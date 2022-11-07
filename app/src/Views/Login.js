import {
	Image,
	Linking,
	StyleSheet,
	Text,
	TouchableOpacity,
	View,
} from 'react-native';
import React from 'react';
import Container from '../Components/Container';
import loginHeader from '../../assets/images/loginHeader.jpg';
import { Colors, Fonts } from '../style';
import bthLogo from '../../assets/images/BTHLogo.png';
import TextField from '../Components/TextField';
import { useState, useRef } from 'react';
import Button from '../Components/Button';
import { t } from '../locale/translate';

const Login = () => {
	const [login, setLogin] = useState('');
	const [password, setPassword] = useState('');
	const [error, setError] = useState(false);

	const passwordRef = useRef();

	const handleLoginChange = (text) => {
		setLogin(text);
		setError(false);
	};

	const handlePasswordChange = (text) => {
		setPassword(text);
		setError(false);
	};

	const handleLogin = () => {
		if (login.length < 1 || password.length < 1) {
			setError(true);
			return;
		}
		console.log('Login: ', login);
		console.log('Password: ', password);
	};

	const goToRestartPassword = () => {
		Linking.openURL('https://personalkonto.bth.se/employee/employee-start');
	};

	return (
		<Container style={styles.container}>
			<View style={styles.header}>
				<Image source={loginHeader} style={styles.image} />
				<View style={styles.colorOverlay} />
				<View style={styles.logoContainer}>
					<Image source={bthLogo} style={styles.logo} />
				</View>
			</View>
			<View style={styles.content}>
				<Text style={styles.title}>{t('login')}</Text>
				<View style={styles.form}>
					<View style={styles.input}>
						<Text
							style={[
								styles.label,
								{
									color: error
										? Colors.danger.regular
										: Colors.secondary.regular,
								},
							]}
						>
							{t('username')}
						</Text>
						<TextField
							value={login}
							onSubmitEditing={() => {
								passwordRef.current.focus();
							}}
							returnKeyType="next"
							placeholder={t('student_acronym_ph')}
							onChangeText={handleLoginChange}
							autoCompleteType="username"
							error={error}
						/>
					</View>
					<View style={styles.input}>
						<Text
							style={[
								styles.label,
								{
									color: error
										? Colors.danger.regular
										: Colors.secondary.regular,
								},
							]}
						>
							{t('password')}
						</Text>
						<TextField
							inputRef={passwordRef}
							value={password}
							placeholder={t('password_ph')}
							returnKeyType="done"
							onChangeText={handlePasswordChange}
							onSubmitEditing={handleLogin}
							autoCompleteType="password"
							secureTextEntry
							error={error}
						/>
					</View>
				</View>
				<TouchableOpacity onPress={goToRestartPassword}>
					<Text style={styles.forgotPassword}>{t('problemes_with_login')}</Text>
				</TouchableOpacity>
				<Button
					style={styles.button}
					onPress={handleLogin}
					textStyle={styles.buttonText}
				>
					{t('login')}
				</Button>
			</View>
		</Container>
	);
};

export default Login;

const styles = StyleSheet.create({
	container: {
		paddingHorizontal: 0,
		justifyContent: 'flex-start',
	},
	header: {
		width: '100%',
		height: '50%',
	},
	image: {
		width: '100%',
		height: '100%',
		resizeMode: 'cover',
	},
	colorOverlay: {
		...StyleSheet.absoluteFillObject,
		backgroundColor: `rgba(5, 80, 100, .3)`,
	},
	logoContainer: {
		...StyleSheet.absoluteFillObject,
		justifyContent: 'center',
		alignItems: 'center',
	},
	logo: {
		width: '65%',
		height: '65%',
		resizeMode: 'contain',
	},
	content: {
		width: '80%',
		justifyContent: 'center',
		alignItems: 'center',
	},
	title: {
		fontSize: Fonts.size.h1,
		marginTop: 20,
	},
	form: {
		width: '100%',
	},
	input: {
		width: '100%',
		marginVertical: 10,
	},
	label: {
		fontSize: Fonts.size.p,
	},
	forgotPassword: {
		fontSize: Fonts.size.p,
		color: Colors.primary.regular,
		marginBottom: 20,
	},
	button: {
		padding: 15,
	},
	buttonText: {
		fontSize: Fonts.size.h3,
		fontFamily: Fonts.Inter_Light,
	},
});
