import { StyleSheet, TextInput, TouchableOpacity, View } from 'react-native';
import React from 'react';
import { Colors, Fonts } from '../style';
import { Ionicons } from '@expo/vector-icons';
const TextField = ({
	style = {},
	onChangeText = () => {},
	value = '',
	placeholder = 'Enter some text...',
	keyboardType = 'default',
	keyboardAppearance = 'default',
	autoCompleteType = 'off',
	secureTextEntry = false,
	canShowPassword = true,
	onSubmitEditing = () => {},
	returnKeyType = 'done',
	inputRef,
	error = false,
}) => {
	const [isHiddenText, setIsHiddenText] = React.useState(secureTextEntry);

	const togglePasswordVisibility = () => {
		if (canShowPassword && secureTextEntry) {
			setIsHiddenText(!isHiddenText);
		}
	};

	const ShowHideIcon = () => {
		return (
			<TouchableOpacity
				style={styles.toggleIcon}
				onPress={togglePasswordVisibility}
			>
				{isHiddenText ? (
					<Ionicons name="eye" size={24} color={Colors.secondary.regular} />
				) : (
					<Ionicons name="eye-off" size={24} color={Colors.secondary.primary} />
				)}
			</TouchableOpacity>
		);
	};

	return (
		<View
			style={[
				styles.container,
				style,
				{ borderWidth: error ? 1 : 0, borderColor: Colors.danger.light },
			]}
		>
			<TextInput
				ref={inputRef}
				onSubmitEditing={onSubmitEditing}
				style={[styles.input, style]}
				placeholder={placeholder}
				onChangeText={onChangeText}
				value={value}
				returnKeyType={returnKeyType}
				keyboardType={keyboardType}
				keyboardAppearance={keyboardAppearance}
				autoCompleteType={autoCompleteType}
				secureTextEntry={isHiddenText}
			/>
			{secureTextEntry && canShowPassword && <ShowHideIcon />}
		</View>
	);
};

export default TextField;

const styles = StyleSheet.create({
	container: {
		backgroundColor: Colors.grey.light,
		padding: 10,
		borderRadius: 10,
		width: '100%',
		marginVertical: 10,
		flexDirection: 'row',
		alignItems: 'center',
		justifyContent: 'space-between',
	},
	input: {
		fontFamily: Fonts.Inter_Regular,
		width: '90%',
	},
});
