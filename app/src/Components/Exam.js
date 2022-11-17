import { StyleSheet, Text, View } from 'react-native';
import React from 'react';
import { t } from '../locale/translate';
import { Colors, Fonts } from '../style';
import Button from './Button';

const Exam = ({ code, name, date, time, room, registered = false }) => {
	const registerExam = () => {};
	const unregisterExam = () => {};

	return (
		<View style={styles.container}>
			<Text style={styles.code}>{code}</Text>
			<Text>{name}</Text>
			<Text>Date: {date}</Text>
			<Text>Time: {time}</Text>
			<Text>Room: {room}</Text>
			<Button
				onPress={registered ? unregisterExam : registerExam}
				type={registered ? 'danger' : 'primary'}
				onCli
				style={styles.buttonStyle}
			>
				{registered ? t('unRegister') : t('register')}
			</Button>
		</View>
	);
};

export default Exam;

const styles = StyleSheet.create({
	container: {
		padding: 10,
		elevation: 10,
		width: '100%',
		backgroundColor: Colors.snowWhite,
		justifyContent: 'flex-start',
		borderRadius: 15,
		margin: 5,
	},
	code: {
		fontSize: Fonts.size.h3,
		alignSelf: 'flex-start',
		fontFamily: Fonts.Inter_Bold,
	},
	buttonStyle: {
		width: '50%',
		alignSelf: 'flex-end',
	},
});
