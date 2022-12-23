import { StyleSheet, Text, View } from 'react-native';
import React from 'react';
import { t } from '../locale/translate';
import { Colors, Fonts } from '../style';
import Button from './Button';
import { unregisterExam } from '../helpers/APIManager';

const Exam = ({
	code,
	name,
	date,
	time,
	room,
	key = 0,
	registered = false,
	HandleRegisterExam = () => {},
	HandleUnregisterExam = () => {},
}) => {
	return (
		<View key={key} style={styles.container}>
			<Text style={styles.code}>{code}</Text>
			<Text>{name}</Text>
			<Text>
				{t('date')}: {date}
			</Text>
			<Text>
				{t('time')}: {time}
			</Text>
			<Text>
				{t('room')}: {room}
			</Text>
			<Button
				onPress={registered ? HandleUnregisterExam : HandleRegisterExam}
				type={registered ? 'danger' : 'primary'}
				onCli
				style={styles.buttonStyle}
			>
				{registered ? t('unregister') : t('register')}
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
