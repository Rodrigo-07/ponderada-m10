// frontend/app/_layout.tsx
import { Stack } from 'expo-router';
import { GestureHandlerRootView } from 'react-native-gesture-handler';

export default function RootLayout() {
  return (
    <GestureHandlerRootView style={{ flex: 1 }}>
      <Stack
        screenOptions={{
          headerStyle: { backgroundColor: '#0a0a23' },
          headerTintColor: '#fff',
        }}
      />
    </GestureHandlerRootView>
  );
}
