import { StyleSheet, Text, View } from "react-native";
import { createStackNavigator } from '@react-navigation/stack';

import React from 'react';
import Main from '../Views/Main';

const MainNavigation = () => {
	const Stack = createStackNavigator();

	return (
		<Stack.Navigator
			initialRouteName="Login"
			screenOptions={{
				headerShown: false,
			}}
		>
			<Stack.Screen name="Login" component={Main} />
		</Stack.Navigator>
	);
};

export default MainNavigation;

const styles = StyleSheet.create({});
