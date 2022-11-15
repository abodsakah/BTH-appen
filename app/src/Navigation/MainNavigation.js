import { StyleSheet, Text, View } from "react-native";
import { createStackNavigator } from '@react-navigation/stack';

import React from 'react';
import Main from '../Views/Main';
import TabBarNavigation from './TabBarNavigation';

const MainNavigation = () => {
	const Stack = createStackNavigator();

	return (
		<Stack.Navigator
			initialRouteName="Main"
			screenOptions={{
				headerShown: false,
			}}
		>
			<Stack.Screen name="Main" component={TabBarNavigation} />
		</Stack.Navigator>
	);
};

export default MainNavigation;

const styles = StyleSheet.create({});
