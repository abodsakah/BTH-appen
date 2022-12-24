import { StyleSheet, View, Text, Settings } from 'react-native'
import React, { useEffect, useState } from 'react';
import Container from '../Components/Container';
import { Colors, Fonts } from '../style';
import {
	AntDesign,
	Ionicons,
	Entypo,
	MaterialCommunityIcons,
} from '@expo/vector-icons';
import OptionContainer from '../Components/OptionContainer';
import { t } from '../locale/translate';
import * as SecureStore from 'expo-secure-store';
import * as Updates from 'expo-updates';
import TriggerContainer from '../Components/TriggerContainer';
import AsyncStorage from '@react-native-async-storage/async-storage';

const Profile = ({ navigation }) => {
	const [user, setUser] = useState(null);
	const [notificationsEnabled, setNotificationsEnabled] = useState(true);

	const navigateToLanguages = () => {
		navigation.navigate('Languages');
	};

	const navigateToAbout = () => {
		navigation.navigate('About');
	};

	const logout = async () => {
		await SecureStore.deleteItemAsync('user');
		Updates.reloadAsync();
	};

	const getUser = async () => {
		const res = await SecureStore.getItemAsync('user');
		if (res) {
			setUser(JSON.parse(res));
		}
	};

	const triggerNotifications = async () => {
		setNotificationsEnabled(!notificationsEnabled);

		if (!notificationsEnabled) {
			await AsyncStorage.setItem('notificationsEnabled', 'false');
		} else {
			await AsyncStorage.setItem('notificationsEnabled', 'true');
		}
	};

	useEffect(() => {
		getUser();
	}, []);

	return (
		<Container style={styles.container}>
			<Text style={styles.heading}>{t('profile')}</Text>
			<OptionContainer
				text={user?.user.name}
				Icon={() => <Ionicons name="md-person" size={30}></Ionicons>}
			/>
			<Text style={styles.heading}>{t('more')}</Text>
			<TriggerContainer
				text={t('Notifications')}
				Icon={() => (
					<Ionicons
						name="notifications-outline"
						size={30}
						color={Colors.primary.regular}
					/>
				)}
				onValueChange={triggerNotifications}
				value={notificationsEnabled}
			/>
			<OptionContainer
				onPress={navigateToLanguages}
				text={t('appLanguage')}
				Icon={() => (
					<Entypo name="language" size={30} color={Colors.primary.regular} />
				)}
			/>
			<OptionContainer
				text={t('logout')}
				Icon={() => (
					<MaterialCommunityIcons
						name="logout-variant"
						size={30}
						color={Colors.primary.regular}
					/>
				)}
				onPress={logout}
			/>
			<OptionContainer
				text={t('about')}
				Icon={() => (
					<Ionicons
						name="md-information-circle"
						size={30}
						color={Colors.primary.regular}
					/>
				)}
				onPress={navigateToAbout}
			/>
		</Container>
	);
};

export default Profile;

const styles = StyleSheet.create({
	container: {
		justifyContent: 'flex-start',
	},
	heading: {
		fontSize: Fonts.size.h2,
		alignSelf: 'flex-start',
		fontFamily: Fonts.Inter_Bold,
		padding: 5,
		margin: 3,
	},
});