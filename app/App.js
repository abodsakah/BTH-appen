import { StatusBar } from 'expo-status-bar';
import * as Device from 'expo-device';
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
import AsyncStorage from '@react-native-async-storage/async-storage';
import { setLanguage, t } from './src/locale/translate';
import { useEffect, useState, useRef } from 'react';
import { Colors } from './src/style';
import * as Notifications from 'expo-notifications';

async function registerForPushNotificationsAsync() {
	let token;
	if (Device.isDevice) {
		const { status: existingStatus } =
			await Notifications.getPermissionsAsync();
		console.log('device');
		let finalStatus = existingStatus;
		if (existingStatus !== 'granted') {
			const { status } = await Notifications.requestPermissionsAsync();
			finalStatus = status;
		}
		if (finalStatus !== 'granted') {
			alert(t('pushTokenGetFailed'));
			return;
		}
		token = (await Notifications.getExpoPushTokenAsync()).data;
		console.log('token: ', token);
	} else {
		alert('Must use physical device for Push Notifications');
	}

	if (Platform.OS === 'android') {
		Notifications.setNotificationChannelAsync('default', {
			name: 'default',
			importance: Notifications.AndroidImportance.MAX,
			vibrationPattern: [0, 250, 250, 250],
			lightColor: Colors.primary.regular,
		});
	}

	return token;
}

export default function App() {
	const [expoPushToken, setExpoPushToken] = useState();
	const [notification, setNotification] = useState(false);

	const notificationListener = useRef();
	const responseListener = useRef();

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

	const getPreferredLanguageAndApply = async () => {
		const lang = await AsyncStorage.getItem('language');
		if (lang) {
			setLanguage(lang);
		}
	};

	useEffect(() => {
		getPreferredLanguageAndApply();

		// Expo push notifications
		registerForPushNotificationsAsync().then((token) =>
			setExpoPushToken(token)
		);

		// This listener is fired whenever a notification is received while the app is foregrounded
		notificationListener.current =
			Notifications.addNotificationReceivedListener((notification) => {
				setNotification(notification);
			});

		// This listener is fired whenever a user taps on or interacts with a notification (works when app is foregrounded, backgrounded, or killed)
		responseListener.current =
			Notifications.addNotificationResponseReceivedListener((response) => {
				console.log(response);
			});

		return () => {
			Notifications.removeNotificationSubscription(notificationListener);
			Notifications.removeNotificationSubscription(responseListener);
		};
	}, []);

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
