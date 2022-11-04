import { StyleSheet, Text, TouchableHighlight, View } from 'react-native';
import React from 'react';
import { Colors, Fonts } from '../style';

const Button = ({
	style = {},
	textStyle = {},
	children,
	type = 'primary',
	onPress = () => {},
}) => {
	return (
		<TouchableHighlight
			activeOpacity={0.5}
			underlayColor={Colors[type].light}
			onPress={onPress}
			style={[
				styles.container,
				style,
				{ backgroundColor: Colors[type].regular },
			]}
		>
			<Text style={[styles.text, textStyle]}>{children}</Text>
		</TouchableHighlight>
	);
};

export default Button;

const styles = StyleSheet.create({
	container: {
		padding: 10,
		borderRadius: 5,
		width: '100%',
		marginVertical: 10,
	},
	text: {
		color: Colors.snowWhite,
		fontFamily: Fonts.light,
	},
});
