import { StyleSheet, Text, View } from "react-native";
import { createStackNavigator } from '@react-navigation/stack';

import React from 'react';
import Main from '../Views/Main';
import Login from '../Views/Login';
import Map from '../Views/Map';

const MainNavigation = () => {
	const Stack = createStackNavigator();

	return (
		<Stack.Navigator
			screenOptions={{
				headerShown: false,
			}}
		>
			<Stack.Screen name="Maps" component={Map} />
			<Stack.Screen name="Login" component={Login} />
		</Stack.Navigator>
	);
};

export default MainNavigation;

const styles = StyleSheet.create({});
