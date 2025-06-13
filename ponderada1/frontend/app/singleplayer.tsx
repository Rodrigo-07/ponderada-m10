import React, { useEffect, useState } from 'react';
import { View, Text, StyleSheet, TouchableOpacity, Alert, ActivityIndicator } from 'react-native';
import Card from '../components/Card';
import { api } from '../hooks/useApi';
import { Stack } from 'expo-router';

type Game = {
  game_id: string;
  drawn_cards: string[];
  card_sum: number;
  result: string;
};

export default function Singleplayer() {
  const [game, setGame] = useState<Game>();
  const [loading, setLoading] = useState(false);

  useEffect(() => { newGame(); }, []);

  async function newGame() {
    try {
      const { data } = await api.post<Game>('/create-game', { player_name: 'Player' });
      setGame(data);
    } catch (err) {
      Alert.alert('Erro', 'Não foi possível iniciar o jogo.');
    }
  }

  async function move(move: 'draw' | 'stop') {
    if (!game) return;
    setLoading(true);
    try {
      const { data } = await api.post('/make-move-singleplayer', { game_id: game.game_id, move });
      setGame(data.game || data);
    } catch (e: any) {
      Alert.alert('Erro', e.response?.data?.error || 'Falha na jogada');
    } finally { setLoading(false); }
  }

  if (!game) return <ActivityIndicator style={{ flex: 1 }} color="#fff" />;

  const finished = game.result !== 'in_progress';

  return (
    <View style={styles.container}>
      <Stack.Screen options={{ title: 'Singleplayer' }} />
      <Text style={styles.text}>Soma: {game.card_sum}</Text>
      <View style={styles.hand}>
        {game.drawn_cards.map((c, i) => <Card key={i} code={c} />)}
      </View>

      {finished && <Text>Resultado: {game.result}</Text>}
      
      {!finished && (
        <View style={styles.row}>
          <Btn title="Comprar" onPress={() => move('draw')} disabled={loading} />
          <Btn title="Parar" onPress={() => move('stop')} disabled={loading} />
        </View>
      )}
      <Btn title="Reiniciar" onPress={newGame} style={{ marginTop: 24 }} />
    </View>
  );
}

function Btn({ title, onPress, disabled, style }: any) {
  return (
    <TouchableOpacity style={[styles.btn, style, disabled && { opacity: 0.5 }]} onPress={onPress} disabled={disabled}>
      <Text style={styles.btnText}>{title}</Text>
    </TouchableOpacity>
  );
}

const styles = StyleSheet.create({
  container: { flex: 1, backgroundColor: '#0a0a23', padding: 16 },
  text: { color: '#fff', fontSize: 18, marginVertical: 8 },
  hand: { flexDirection: 'row', flexWrap: 'wrap', marginVertical: 16 },
  row: { flexDirection: 'row', gap: 12 },
  btn: { backgroundColor: '#1e90ff', padding: 16, borderRadius: 8, flex: 1, alignItems: 'center' },
  btnText: { color: '#fff', fontSize: 16 },
});
