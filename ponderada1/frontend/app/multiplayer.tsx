import React, { useState, useRef } from 'react';
import {
  View, Text, StyleSheet, TouchableOpacity,
  Alert, ActivityIndicator, SafeAreaView, TextInput,
} from 'react-native';
import { LinearGradient } from 'expo-linear-gradient';
import { Stack } from 'expo-router';
import Card from '../components/Card';
import { api } from '../hooks/useApi';

// api typing
type Game = {
  game_id: string;
  player1_name: string;
  player2_name: string;
  player1_visible: string[];
  player2_visible: string[];
  player1_hidden: string;
  player2_hidden: string;
  player1_extra?: string[];
  player2_extra?: string[];
  player1_score: number;
  player2_score: number;
  result: 'in_progress' | 'player1' | 'player2' | 'draw';
  current_turn: 'player1' | 'player2';
  round: number;
};

const ScaledCard = ({
  code, hidden, scale = 1,
}: { code?: string; hidden?: boolean; scale?: number }) => {
  if (!code && !hidden) return null;
  return (
    <View style={{ transform: [{ scale }], marginRight: 8, marginBottom: 8 }}>
      <Card code={code} hidden={hidden} />
    </View>
  );
};

const Btn = ({ title, onPress, disabled, style }: any) => (
  <TouchableOpacity
    style={[styles.btn, style, disabled && styles.btnDisabled]}
    onPress={onPress}
    disabled={disabled}
    activeOpacity={0.85}
  >
    <Text style={styles.btnText}>{title}</Text>
  </TouchableOpacity>
);

export default function Multiplayer() {
  const [name1, setName1] = useState('');
  const [name2, setName2] = useState('');
  const [setupDone, setSetupDone] = useState(false);

  const [game, setGame] = useState<Game>();
  const [loading, setLoading] = useState(false);

  const [overlay, setOverlay] = useState<'none' | 'preview' | 'handoff'>('none');
  const [previewHand, setPreviewHand] = useState<string[]>([]);

  async function newGame(p1: string, p2: string) {
    try {
      const { data } = await api.post<Game>('/create-game-multiplayer', {
        player1_name: p1,
        player2_name: p2,
      });
      setGame(data);
      setOverlay('none');
    } catch {
      Alert.alert('Erro', 'Falha ao criar partida');
    }
  }

  async function play(move: 'draw' | 'pass') {
    if (!game || loading || overlay !== 'none') return;
    setLoading(true);
    try {
      const { data } = await api.post<Game>('/make-move-multiplayer', {
        game_id: game.game_id,
        player_name: game.current_turn === 'player1'
          ? game.player1_name
          : game.player2_name,
        move,
      });

      const isP1 = data.current_turn === 'player2';
      const myVis = isP1 ? data.player1_visible : data.player2_visible;
      const myHid = isP1 ? data.player1_hidden : data.player2_hidden;
      const myExt = isP1 ? data.player1_extra : data.player2_extra;
      setPreviewHand([...myVis, myHid, ...(myExt ?? [])].filter(Boolean));

      setGame(data);
      if (data.result === 'in_progress') setOverlay('preview');
    } catch (e: any) {
      Alert.alert('Erro', e.response?.data?.error || 'Jogada inválida');
    } finally {
      setLoading(false);
    }
  }

  if (!setupDone) {
    return (
      <LinearGradient colors={['#173d2b', '#0d241a']} style={styles.full}>
        <SafeAreaView style={styles.fullCentered}>
          <Text style={styles.setupTitle}>Nomes dos Jogadores</Text>
          <TextInput
            style={styles.input}
            placeholder="Jogador 1"
            placeholderTextColor="#aaa"
            value={name1}
            onChangeText={setName1}
          />
          <TextInput
            style={styles.input}
            placeholder="Jogador 2"
            placeholderTextColor="#aaa"
            value={name2}
            onChangeText={setName2}
          />
          <Btn
            title="Iniciar Partida"
            style={{ marginTop: 20 }}
            disabled={!name1 || !name2}
            onPress={() => {
              setSetupDone(true);
              newGame(name1, name2);
            }}
          />
        </SafeAreaView>
      </LinearGradient>
    );
  }

  if (!game) {
    return (
      <LinearGradient colors={['#173d2b', '#0d241a']} style={styles.full}>
        <ActivityIndicator size="large" color="#fff" />
      </LinearGradient>
    );
  }

  const p1Turn = game.current_turn === 'player1';

  const meName = p1Turn ? game.player1_name : game.player2_name;
  const oppName = p1Turn ? game.player2_name : game.player1_name;

  const meVis = p1Turn ? game.player1_visible : game.player2_visible;
  const meHid = p1Turn ? game.player1_hidden : game.player2_hidden;
  const meExtra = p1Turn ? game.player1_extra : game.player2_extra;

  const oppVis = p1Turn ? game.player2_visible : game.player1_visible;
  const oppHid = p1Turn ? game.player2_hidden : game.player1_hidden;
  const oppExtra = p1Turn ? game.player2_extra : game.player1_extra;

  const meScore = p1Turn ? game.player1_score : game.player2_score;
  const busted = meScore > 21;
  const finished = game.result === 'player1' || game.result === 'player2' || game.result === 'draw';

  const myHandNow = [...meVis, meHid, ...(meExtra ?? [])].filter(Boolean);

  return (
    <LinearGradient colors={['#173d2b', '#0d241a']} style={styles.full}>
      <SafeAreaView style={styles.full}>
        <Stack.Screen options={{ title: 'Blackjack' }} />

        <View style={styles.scoreboard}>
          <Text style={styles.scoreTitle}>{meName}</Text>
          <Text style={styles.myScore}>{meScore}{busted && ' ✖'}</Text>
          <Text style={styles.round}>Rodada {game.round}</Text>
        </View>

        <Text style={styles.section}>Suas cartas {busted && '(Estourei)'}</Text>
        <View style={styles.hand}>
          {myHandNow.map((c, i) => (
            <ScaledCard key={`mc-${i}`} code={c} scale={1.2} />
          ))}
        </View>

        <Text style={styles.section}>Cartas de {oppName}</Text>
        <View style={styles.hand}>
          {oppVis.map((c, i) => (
            <ScaledCard key={`ov-${i}`} code={c} scale={0.8} />
          ))}
          <ScaledCard code={oppHid} hidden={!finished} scale={0.8} />
          {(oppExtra ?? []).map((c, i) => (
            <ScaledCard
              key={`ox-${i}`}
              code={c}
              hidden={!finished}
              scale={0.8}
            />
          ))}
        </View>

        {!finished && (
          <View style={styles.actions}>
            <Btn
              title="Comprar"
              onPress={() => play('draw')}
              disabled={loading || busted || overlay !== 'none'}
            />
            <Btn
              title="Parar"
              onPress={() => play('pass')}
              disabled={loading || overlay !== 'none'}
            />
          </View>
        )}

        {finished && (
          <View style={styles.resultBox}>
            <Text style={styles.resultText}>
              {game.result === 'draw'
                ? 'EMPATE'
                : `${game.result === 'player1' ? game.player1_name : game.player2_name} VENCEU`}
            </Text>
            <Btn
              title="Nova Partida"
              onPress={() => newGame(name1, name2)}
              style={{ marginTop: 8 }}
            />
          </View>
        )}

        {overlay === 'preview' && (
          <View style={styles.overlay}>
            <Text style={styles.overlayText}>Suas cartas</Text>
            <View style={styles.hand}>
              {previewHand.map((c, i) => (
                <ScaledCard key={`ph-${i}`} code={c} />
              ))}
            </View>
            <Btn
              title="Próxima"
              style={{ marginTop: 24 }}
              onPress={() => setOverlay('handoff')}
            />
          </View>
        )}

        {overlay === 'handoff' && (
          <View style={styles.overlay}>
            <Text style={styles.overlayText}>Passe o celular para {meName}</Text>
            <Btn
              title="Pronto"
              style={{ marginTop: 24 }}
              onPress={() => setOverlay('none')}
            />
          </View>
        )}
      </SafeAreaView>
    </LinearGradient>
  );
}

const styles = StyleSheet.create({
  full: { flex: 1 },
  fullCentered: { flex: 1, justifyContent: 'center', alignItems: 'center', padding: 24 },
  setupTitle: { color: '#fff', fontSize: 22, marginBottom: 24 },
  input: {
    width: '80%', backgroundColor: '#fff', borderRadius: 8,
    paddingHorizontal: 12, paddingVertical: 10, fontSize: 16, marginVertical: 6,
  },
  scoreboard: { alignItems: 'center', marginBottom: 12 },
  scoreTitle: { color: '#e8e8e8', fontSize: 18 },
  myScore: { color: '#fff', fontSize: 32, fontWeight: 'bold', marginVertical: 2 },
  round: { color: '#ffd24c', fontSize: 16 },
  section: { color: '#c8c8c8', fontSize: 16, marginTop: 8, marginLeft: 8 },
  hand: { flexDirection: 'row', flexWrap: 'wrap', marginLeft: 8 },
  actions: {
    flexDirection: 'row', justifyContent: 'space-evenly',
    marginTop: 12, paddingHorizontal: 8,
  },
  btn: {
    backgroundColor: '#ffb400',
    paddingVertical: 14, paddingHorizontal: 24,
    borderRadius: 10, shadowColor: '#000', shadowOpacity: 0.4, shadowRadius: 4,
    shadowOffset: { width: 0, height: 2 }, elevation: 4, alignItems: 'center',
  },
  btnDisabled: { opacity: 0.55 },
  btnText: { color: '#322200', fontWeight: 'bold', fontSize: 16 },
  resultBox: { alignItems: 'center', marginTop: 24 },
  resultText: { color: '#fff', fontSize: 26, fontWeight: 'bold' },
  overlay: {
    position: 'absolute', top: 0, left: 0, right: 0, bottom: 0,
    backgroundColor: 'rgba(0, 0, 0, 1)', padding: 32,
    justifyContent: 'center', alignItems: 'center',
  },
  overlayText: { color: '#fff', fontSize: 20, textAlign: 'center', marginBottom: 16 },
});
