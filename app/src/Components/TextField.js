import { StyleSheet, Text, TextInput, View } from 'react-native';
import React from 'react';
import { Colors, Fonts } from '../style';

const TextField = ({
	style = {},
	onChangeText = () => {},
	value = '',
	placeholder = 'Enter some text...',
	keyboardType = 'default',
	keyboardAppearance = 'default',
}) => {
	return (
		<TextInput
			style={[styles.input, style]}
			placeholder={placeholder}
			onChangeText={onChangeText}
			value={value}
			keyboardType={keyboardType}
			keyboardAppearance={keyboardAppearance}
		/>
	);
};

export default TextField;

const styles = StyleSheet.create({
	input: {
		backgroundColor: Colors.grey.light,
		padding: 10,
		borderRadius: 10,
		fontFamily: Fonts.Inter_Regular,
		width: '100%',
		marginVertical: 10,
	},
});
