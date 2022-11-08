import { StatusBar } from 'expo-status-bar';
import { ActivityIndicator, StyleSheet, Text, View } from 'react-native';
import Button from './src/Components/Button';
import Container from './src/Components/Container';
import TextField from './src/Components/TextField';
import {
	Inter_100Thin,
	Inter_200ExtraLight,
	Inter_300Light,
	Inter_400Regular,
	Inter_500Medium,
	Inter_600SemiBold,
	Inter_700Bold,
	Inter_800ExtraBold,
	Inter_900Black,
	useFonts,
} from '@expo-google-fonts/inter';
import Login from './src/Views/Login';
import MainNavigation from './src/Navigation/MainNavigation';
import { NavigationContainer } from '@react-navigation/native';

export default function App() {
	let [fontsLoaded] = useFonts({
		Inter_100Thin,
		Inter_200ExtraLight,
		Inter_300Light,
		Inter_400Regular,
		Inter_500Medium,
		Inter_600SemiBold,
		Inter_700Bold,
		Inter_800ExtraBold,
		Inter_900Black,
	});

	if (!fontsLoaded) {
		return <ActivityIndicator />;
	}

	return (
		<NavigationContainer>
			<MainNavigation />
		</NavigationContainer>
	);
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    alignItems: 'center',
    justifyContent: 'center',
  },
});
