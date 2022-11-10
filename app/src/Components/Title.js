import { StyleSheet, Text, View } from 'react-native';
import React from 'react';
import { Colors, Fonts } from '../style';

const Title = ({ children, style }) => {
	return <Text style={[styles.title, style]}>{children}</Text>;
};

export default Title;

const styles = StyleSheet.create({
	title: {
		fontFamily: Fonts.Inter_ExtraBold,
		fontSize: Fonts.size.h1,
		color: Colors.secondary,
	},
});
