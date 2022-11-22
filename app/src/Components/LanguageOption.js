import { StyleSheet, Text, TouchableOpacity, View } from 'react-native';
import React from 'react';
import { Colors, Fonts } from '../style';

const LanguageOption = ({ name, selected, isRTL, onPress = () => {} }) => {
	return (
		<TouchableOpacity onPress={onPress} style={styles.container}>
			<View style={styles.selectButton}>
				{selected && <View style={styles.selected} />}
			</View>
			<Text style={[styles.name, { textAlign: isRTL ? 'right' : 'left' }]}>
				{name}
			</Text>
		</TouchableOpacity>
	);
};

export default LanguageOption;

const styles = StyleSheet.create({
	container: {
		flexDirection: 'row',
		alignItems: 'center',
		marginVertical: 10,
		elevation: 10,
		padding: 10,
		width: '100%',
		height: 50,
		backgroundColor: Colors.snowWhite,
		borderRadius: 15,
	},
	selectButton: {
		width: 30,
		height: 30,
		borderRadius: 15,
		borderWidth: 3,
		borderColor: Colors.primary.regular,
		marginRight: 10,
		alignItems: 'center',
		justifyContent: 'center',
	},
	selected: {
		width: 15,
		height: 15,
		borderRadius: 15,
		backgroundColor: Colors.primary.regular,
	},
	name: {
		fontSize: Fonts.size.h3,
	},
});
