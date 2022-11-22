import { StyleSheet, Text, View } from 'react-native';
import React from 'react';
import { Entypo } from '@expo/vector-icons';
import { Colors } from '../style';
import { TouchableRipple } from 'react-native-paper';

const OptionContainer = ({
	text = '',
	Icon = () => {},
	onPress = () => {},
}) => {
	return (
		<View style={styles.container}>
			<TouchableRipple borderless onPress={onPress} style={styles.rippleEffect}>
				<>
					<Icon />
					<Text>{text}</Text>
					<Entypo
						name="chevron-thin-right"
						size={24}
						color={Colors.primary.regular}
					/>
				</>
			</TouchableRipple>
		</View>
	);
};

export default OptionContainer;

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
	},
	rippleEffect: {
		borderRadius: 10,
		padding: 15,
		height: '100%',
		marginVertical: 5,
		backgroundColor: Colors.snowWhite,
		width: '100%',
		alignItems: 'center',
		flexDirection: 'row',
		justifyContent: 'space-between',
	},
});
