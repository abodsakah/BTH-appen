import { StyleSheet, Text, View } from 'react-native';
import React from 'react';
import Container from '../Components/Container';
import { Fonts } from '../style';
import { t } from '../locale/translate';

const About = () => {
	return (
		<Container>
			<Text style={styles.title}>About</Text>
			<View style={styles.textContainer}>
				<Text>
					{t('aboutText')}
					{'\n'}
					{'\n'}
				</Text>
				<Text>
					{t('developedBy')}
					{'\n'}- Abdulrahman Sakah{'\n'}- Gabriel Ivarsson{'\n'}- Najem Hamo
					{'\n'}- Noah Håkansson{'\n'}- Sara Madwar
				</Text>
			</View>
		</Container>
	);
};

export default About;

const styles = StyleSheet.create({
	title: {
		fontSize: Fonts.size.h1,
		alignSelf: 'flex-start',
		fontFamily: Fonts.Inter_Bold,
		marginLeft: 10,
	},
	textContainer: {
		flex: 1,
		marginTop: 20,
		marginLeft: 10,
		justifyContent: 'flex-start',
		alignItems: 'flex-start',
	},
});
