import { StyleSheet, Text, View } from 'react-native';
import React from 'react';
import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import Main from '../Views/Main';
import { Ionicons } from '@expo/vector-icons';
import { Colors } from '../style';
import TabBarIcon from '../Components/TabBarIcon';
import { t } from '../locale/translate';
import Exams from '../Views/Exams';
import Profile from '../Views/Profile';
import Map from '../Views/Map';

const TabBarNavigation = () => {
	const Tab = createBottomTabNavigator();

	return (
		<Tab.Navigator
			initialRouteName="Main"
			screenOptions={{
				headerShown: false,
				tabBarStyle: styles.tabBar,
			}}
		>
			<Tab.Screen
				name="News"
				component={Main}
				options={{
					tabBarShowLabel: false,
					tabBarIcon: ({ focused }) => (
						<TabBarIcon
							title={t('news')}
							isFocused={focused}
							Icon={({ color, size }) => (
								<Ionicons name="newspaper-outline" size={size} color={color} />
							)}
						/>
					),
				}}
			/>
			<Tab.Screen
				name="Map"
				component={Map}
				options={{
					tabBarShowLabel: false,
					tabBarIcon: ({ focused }) => (
						<TabBarIcon
							title={t('map')}
							isFocused={focused}
							Icon={({ color, size }) => (
								<Ionicons name="ios-map-outline" size={size} color={color} />
							)}
						/>
					),
				}}
			/>
			<Tab.Screen
				name="Exams"
				component={Exams}
				options={{
					tabBarShowLabel: false,
					tabBarIcon: ({ focused }) => (
						<TabBarIcon
							title={t('exams')}
							isFocused={focused}
							Icon={({ color, size }) => (
								<Ionicons
									name="md-document-text-outline"
									size={size}
									color={color}
								/>
							)}
						/>
					),
				}}
			/>
			<Tab.Screen
				name="Profile"
				component={Profile}
				options={{
					tabBarShowLabel: false,
					tabBarIcon: ({ focused }) => (
						<TabBarIcon
							title={t('profile')}
							isFocused={focused}
							Icon={({ color, size }) => (
								<Ionicons name="person-outline" size={size} color={color} />
							)}
						/>
					),
				}}
			/>
		</Tab.Navigator>
	);
};

export default TabBarNavigation;

const styles = StyleSheet.create({
	tabBar: {
		backgroundColor: Colors.primary.regular,
		height: '10%',
		borderTopLeftRadius: 20,
		borderTopRightRadius: 20,
	},
});
