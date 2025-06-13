import { View, Text, StyleSheet, Image } from 'react-native';

type Props = { code?: string; hidden?: boolean };

export default function Card({ code, hidden }: Props) {
  if (hidden) {
    return (
      <Image
        source={{ uri: 'https://deckofcardsapi.com/static/img/back.png' }}
        style={styles.img}
      />
    );
  }

  return (
    <View style={styles.card}>
      <Image
        source={{ uri: `https://deckofcardsapi.com/static/img/${code}.png` }}
        style={styles.img}
        resizeMode="contain"
      />
    </View>
  );
}

const styles = StyleSheet.create({
  card: {
    width: 60,
    height: 90,
    borderRadius: 6,
    borderWidth: 1,
    borderColor: '#fff',
    backgroundColor: '#2e2e2e',
    justifyContent: 'center',
    alignItems: 'center',
    margin: 4,
  },
  img: { width: 54, height: 82 },        // padding interno de 3 px
  hidden: { backgroundColor: '#555' },
  q: { color: '#fff', fontSize: 28 },
});
