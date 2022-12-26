import { StyleSheet, Text, View } from 'react-native';
import React from 'react';
import { Entypo } from '@expo/vector-icons';
import { Colors } from '../style';
import { Switch, TouchableRipple } from 'react-native-paper';

const TriggerContainer = ({
	text = '',
	Icon = () => {},
	onValueChange = () => {},
	value = false,
}) => {
	return (
		<View style={styles.container}>
			<Icon />
			<Text>{text}</Text>
			<Switch
				value={value}
				onValueChange={onValueChange}
				color={Colors.primary.regular}
			/>
		</View>
	);
};

export default TriggerContainer;

const styles = StyleSheet.create({
	container: {
		borderRadius: 10,
		elevation: 2,
		marginVertical: 5,
		backgroundColor: Colors.snowWhite,
		width: '100%',
		alignItems: 'center',
		flexDirection: 'row',
		justifyContent: 'space-between',
		padding: 15,
	},
});
