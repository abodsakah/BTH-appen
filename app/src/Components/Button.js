import { StyleSheet, Text } from 'react-native';
import React from 'react';
import { Colors, Fonts } from '../style';
import { TouchableRipple } from 'react-native-paper';

const Button = ({
	style = {},
	textStyle = {},
	children,
	type = 'primary',
	onPress = () => {},
	disabled = false,
}) => {
	return (
		<TouchableRipple
			onPress={onPress}
			disabled={disabled}
			style={[
				styles.container,
				style,
				{ backgroundColor: Colors[type].regular, opacity: disabled ? 0.5 : 1 },
			]}
		>
			<Text style={[styles.text, textStyle]}>{children}</Text>
		</TouchableRipple>
	);
};

export default Button;

const styles = StyleSheet.create({
	container: {
		padding: 10,
		borderRadius: 10,
		width: '100%',
		marginVertical: 10,
		justifyContent: 'center',
		alignItems: 'center',
	},
	text: {
		color: Colors.snowWhite,
		fontFamily: Fonts.light,
	},
});
