import { StyleSheet, Text, TouchableOpacity, View } from 'react-native';
import React from 'react';
import Container from '../Components/Container';
import { Colors, Fonts } from '../style';
import { Entypo } from '@expo/vector-icons';
import { useEffect } from 'react';
import AsyncStorage from '@react-native-async-storage/async-storage';
import LanguageOption from '../Components/LanguageOption';
import { t } from '../locale/translate';
import * as Updates from 'expo-updates';

const Languages = ({ navigation }) => {
	const [prefLang, setPrefLang] = React.useState('en');

	const getLanguagePref = async () => {
		const lang = await AsyncStorage.getItem('language');
		if (lang) {
			setPrefLang(lang);
		}
	};

	const setLanguagePref = async (lang) => {
		await AsyncStorage.setItem('language', lang);
		setPrefLang(lang);
		Updates.reloadAsync();
	};

	useEffect(() => {
		getLanguagePref();
	}, []);

	return (
		<Container style={styles.container}>
			<View style={styles.header}>
				<TouchableOpacity
					style={styles.backBtn}
					onPress={() => navigation.goBack()}
				>
					<Entypo name="chevron-left" size={24} color="black" />
				</TouchableOpacity>
				<Text style={styles.title}>{t('appLanguage')}</Text>
			</View>

			<View style={styles.languages}>
				<LanguageOption
					onPress={() => setLanguagePref('en')}
					name="English"
					selected={prefLang === 'en'}
				/>
				<LanguageOption
					onPress={() => setLanguagePref('sv')}
					name="Svenska"
					selected={prefLang === 'sv'}
				/>
				<LanguageOption
					onPress={() => setLanguagePref('ar')}
					name="العربية"
					selected={prefLang === 'ar'}
					isRTL
				/>
			</View>
		</Container>
	);
};

export default Languages;

const styles = StyleSheet.create({
	container: {
		justifyContent: 'flex-start',
	},
	header: {
		flexDirection: 'row',
		alignItems: 'center',
		alignSelf: 'flex-start',
	},
	backBtn: {
		padding: 10,
		width: 50,
	},
	title: {
		fontSize: Fonts.size.h1,
		fontFamily: Fonts.Inter_Bold,
		color: Colors.secondary.regular,
		alignSelf: 'flex-start',
	},
	languages: {
		width: '100%',
	},
});
