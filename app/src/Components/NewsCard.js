import { View, StyleSheet, Text, SafeAreaView, Image} from "react-native";
import React from 'react';
import { Colors, Fonts } from '../style';
import { TouchableRipple } from 'react-native-paper';

const NewsCard = ({ style, title, description, date, onPress = () => {} }) => {
	return (
		<View style={[styles.container, style]}>
			<TouchableRipple onPress={onPress} style={styles.touchableRipple}>
				<View style={{ padding: 20 }}>
					{/* Title */}
					<Text style={styles.title}>{title}</Text>

					{/* Description */}
					<Text style={styles.description}>{description}</Text>

					<View style={styles.data}>
						<Text style={styles.date}>Published: {date}</Text>
					</View>
				</View>
			</TouchableRipple>
		</View>
	);
};

export default NewsCard;

const styles = StyleSheet.create({
	container: {
		width: '100%',
		alignSelf: 'center',
		borderRadius: 20,
		shadowOpacity: 0.5,
		shadowColor: '#000',
		shadowOffset: { width: 5, height: 5 },
		backgroundColor: Colors.snowWhite,
		marginTop: 20,
		elevation: 5,
	},
	touchableRipple: {
		width: '100%',
		borderRadius: 20,
	},
	title: {
		fontSize: Fonts.size.h3,
		fontFamily: Fonts.Inter_Bold,
		merginTop: 10,
	},
	description: {
		fontSize: Fonts.size.h6,
		fontFamily: Fonts.Inter_Medium,
		merginTop: 10,
	},
	data: {
		flexDirection: 'row',
		justifyContent: 'flex-end',
		paddingTop: 10,
	},
	date: {
		fontFamily: Fonts.Inter_Light,
	},
});