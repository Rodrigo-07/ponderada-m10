import { Link } from 'expo-router';
import { StyleSheet, Text, TouchableOpacity, View } from 'react-native';

export default function Home() {
  return (
    <View style={[styles.container, { backgroundColor: '#173d2b' }]}>
      <Text style={styles.title}>Escolha o modo</Text>

      <Link href="/singleplayer" asChild>
        <TouchableOpacity style={styles.button}>
          <Text style={styles.btnText}>Singleplayer</Text>
        </TouchableOpacity>
      </Link>

      <Link href="/multiplayer" asChild>
        <TouchableOpacity style={styles.button}>
          <Text style={styles.btnText}>Multiplayer (Local)</Text>
        </TouchableOpacity>
      </Link>
    </View>
  );
}

const styles = StyleSheet.create({
  container: { flex: 1, backgroundColor: '#0a0a23', alignItems: 'center', justifyContent: 'center', gap: 24 },
  title: { fontSize: 24, color: '#fff' },
  button: { backgroundColor: '#1e90ff', paddingHorizontal: 32, paddingVertical: 16, borderRadius: 8, width: 220, alignItems: 'center' },
  btnText: { color: '#fff', fontSize: 18 },
  btn: {
    backgroundColor: '#ffb400',
  }
});
