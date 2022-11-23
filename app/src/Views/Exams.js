import { ScrollView, StyleSheet, Text, View } from 'react-native';
import React from 'react';
import Container from '../Components/Container';
import Button from '../Components/Button';
import { Colors, Fonts } from '../style';
import Exam from '../Components/Exam';
import { t } from '../locale/translate';

const Exams = () => {
	return (
		<Container style={styles.container}>
			<Text style={styles.title}>{t('exams')}</Text>
			<ScrollView
				style={styles.scrollContainer}
				contentContainerStyle={styles.scrollContentStyle}
			>
				<Text style={styles.heading}>{t('registered_exams')}</Text>
				<Exam
					code="PA1469"
					name="Applikation Utveckling"
					date="2023-05-31"
					time="13:00"
					room="J1240"
					registered
				/>
				<Text style={styles.heading}>{t('coming_exams')}</Text>
				<Exam
					code="DV1628"
					name="Webbutveckling"
					date="2023-05-31"
					time="13:00"
					room="J1240"
				/>
				<Exam
					code="DV1628"
					name="Webbutveckling"
					date="2023-05-31"
					time="13:00"
					room="J1240"
				/>
				<Exam
					code="DV1628"
					name="Webbutveckling"
					date="2023-05-31"
					time="13:00"
					room="J1240"
				/>
				<Exam
					code="DV1628"
					name="Webbutveckling"
					date="2023-05-31"
					time="13:00"
					room="J1240"
				/>
				<Exam
					code="DV1628"
					name="Webbutveckling"
					date="2023-05-31"
					time="13:00"
					room="J1240"
				/>
			</ScrollView>
		</Container>
	);
};

export default Exams;

const styles = StyleSheet.create({
	container: {
		paddingHorizontal: 0,
	},
	title: {
		fontSize: Fonts.size.h1,
		alignSelf: 'flex-start',
		fontFamily: Fonts.Inter_Bold,
		marginLeft: 10,
	},
	scrollContainer: {
		width: '100%',
	},
	scrollContentStyle: {
		paddingHorizontal: 15,
		alignItems: 'center',
	},
	heading: {
		fontSize: Fonts.size.h3,
		margin: 5,
		alignSelf: 'flex-start',
	},
});
