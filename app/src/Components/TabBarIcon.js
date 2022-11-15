import { StyleSheet, Text, View } from 'react-native';
import React from 'react';
import { Colors, Fonts } from '../style';

const TabBarIcon = ({ title, Icon, isFocused }) => {
	return (
		<View style={styles.container}>
			<View
				style={[
					styles.iconContainer,
					{ backgroundColor: isFocused ? Colors.primary.dark : 'transparent' },
				]}
			>
				<Icon color={Colors.snowWhite} size={24} />
			</View>
			<Text style={styles.label}>{title}</Text>
		</View>
	);
};

export default TabBarIcon;

const styles = StyleSheet.create({
	container: {
		alignItems: 'center',
		justifyContent: 'center',
		height: '100%',
	},
	iconContainer: {
		width: 70,
		alignItems: 'center',
		justifyContent: 'center',
		borderRadius: 100,
		padding: 5,
	},
	label: {
		color: Colors.snowWhite,
		fontFamily: Fonts.Inter_Light,
	},
});
