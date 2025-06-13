import React, { useEffect, useState } from 'react';
import {
  View, Text, StyleSheet, TouchableOpacity,
  Alert, ActivityIndicator, SafeAreaView,
} from 'react-native';
import { LinearGradient } from 'expo-linear-gradient';
import { Stack } from 'expo-router';
import Card from '../components/Card';
import { api } from '../hooks/useApi';

type Game = {
  game_id: string;
  drawn_cards: string[];
  card_sum: number;
  result: 'in_progress' | 'win' | 'lose' | 'draw';
};

export default function Singleplayer() {
  const [game, setGame] = useState<Game>();
  const [loading, setLoading] = useState(false);

  useEffect(() => { newGame(); }, []);

  async function newGame() {
    try {
      const { data } = await api.post<Game>('/create-game', { player_name: 'Player' });
      setGame(data);
    } catch {
      Alert.alert('Erro', 'Não foi possível iniciar o jogo.');
    }
  }

  async function move(move: 'draw' | 'stop') {
    if (!game) return;
    setLoading(true);
    try {
      const { data } = await api.post('/make-move-singleplayer', {
        game_id: game.game_id, move
      });
      setGame(data.game || data);
    } catch (e: any) {
      Alert.alert('Erro', e.response?.data?.error || 'Falha na jogada');
    } finally { setLoading(false); }
  }

  if (!game) {
    return (
      <LinearGradient colors={['#173d2b', '#0d241a']} style={styles.full}>
        <ActivityIndicator size="large" color="#fff" />
      </LinearGradient>
    );
  }

  const finished = game.result !== 'in_progress';
  const busted = game.card_sum > 21;

  return (
    <LinearGradient colors={['#173d2b', '#0d241a']} style={styles.full}>
      <SafeAreaView style={styles.full}>
        <Stack.Screen options={{ title: 'Blackjack Solo' }} />

        <View style={styles.scoreboard}>
          <Text style={styles.scoreTitle}>Você</Text>
          <Text style={styles.myScore}>{game.card_sum}{busted && ' ✖'}</Text>
        </View>

        <Text style={styles.section}>Suas cartas</Text>
        <View style={styles.hand}>
          {game.drawn_cards.map((c, i) => <Card key={i} code={c} />)}
        </View>

        {finished && (
          <View style={styles.resultBox}>
            <Text style={styles.resultText}>
              {game.result === 'win' && 'VOCÊ VENCEU'}
              {game.result === 'lose' && 'DERROTA'}
              {game.result === 'draw' && 'EMPATE'}
            </Text>
          </View>
        )}

        {!finished && (
          <View style={styles.actions}>
            <Btn title="Comprar" onPress={() => move('draw')} disabled={loading || busted} />
            <Btn title="Parar" onPress={() => move('stop')} disabled={loading} />
          </View>
        )}

        <Btn title="Nova Partida" onPress={newGame} style={{ marginTop: 24 }} />
      </SafeAreaView>
    </LinearGradient>
  );
}

function Btn({ title, onPress, disabled, style }: any) {
  return (
    <TouchableOpacity
      style={[styles.btn, style, disabled && styles.btnDisabled]}
      onPress={onPress}
      disabled={disabled}
      activeOpacity={0.85}
    >
      <Text style={styles.btnText}>{title}</Text>
    </TouchableOpacity>
  );
}

const styles = StyleSheet.create({
  full: { flex: 1 },
  scoreboard: { alignItems: 'center', marginBottom: 12 },
  scoreTitle: { color: '#e8e8e8', fontSize: 18 },
  myScore: { color: '#fff', fontSize: 32, fontWeight: 'bold', marginVertical: 2 },
  section: { color: '#c8c8c8', fontSize: 16, marginTop: 8, marginLeft: 8 },
  hand: { flexDirection: 'row', flexWrap: 'wrap', marginLeft: 8, marginBottom: 12 },
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
  resultBox: { alignItems: 'center', marginTop: 16 },
  resultText: { color: '#fff', fontSize: 24, fontWeight: 'bold' },
});
