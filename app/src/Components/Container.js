import { StyleSheet, Text, View } from 'react-native';
import React from 'react';
import { SafeAreaView } from 'react-native-safe-area-context';
import { Colors } from '../style';
import { StatusBar } from 'expo-status-bar';

const Container = ({ children, style }) => {
	return (
		<SafeAreaView style={[style, styles.container]}>
			{children}
			<StatusBar style="auto" />
		</SafeAreaView>
	);
};

export default Container;

const styles = StyleSheet.create({
	container: {
		flex: 1,
		backgroundColor: Colors.snowWhite,
		alignItems: 'center',
		justifyContent: 'center',
		paddingHorizontal: 20,
	},
});
