import {
	Dimensions,
	StyleSheet,
	Text,
	TextInput,
	TouchableOpacity,
	View,
} from 'react-native';
import React from 'react';
import Container from '../Components/Container';
import MapView from 'react-native-maps';
import Title from '../Components/Title';
import TextField from '../Components/TextField';
import { AntDesign } from '@expo/vector-icons';
import { Colors } from '../style';
import { LinearGradient } from 'expo-linear-gradient';
import { Buildings } from '../helpers/Constants';

const Map = () => {
	const [search, setSearch] = React.useState('');

	const onSearch = (t) => {
		const searchString = t.toUpperCase();

		const regex = /^[A-Z][1-5][0-9]{2,4}$/;
		if (regex.test(searchString)) {
			const building = searchString[0];
			const floor = Number(searchString[1]);
			const room = searchString.substring(2);

			// TODO: edge cases m, lövsalen, biblioteket

			if (!(building.toLowerCase() in Buildings)) {
				console.log('Building not found');
			}
		} else {
			// show error message
		}
	};

	return (
		<Container style={styles.container}>
			<Title style={styles.title}>Campus Map</Title>
			<View style={styles.mapContainer}>
				<View style={styles.searchFieldContainer}>
					<View style={styles.searchField}>
						<TextInput
							style={styles.searchInput}
							onChangeText={onSearch}
							placeholder="Search"
						/>
						<TouchableOpacity style={styles.searchButton}>
							<AntDesign name="search1" size={24} color={Colors.snowWhite} />
						</TouchableOpacity>
					</View>
				</View>
				<MapView
					style={styles.map}
					initialRegion={{
						latitude: 56.181339,
						longitude: 15.591412,
						latitudeDelta: 0.0922,
						longitudeDelta: 0.0421,
					}}
					initialCamera={{
						center: {
							latitude: 56.181339,
							longitude: 15.591412,
						},
						pitch: 0,
						heading: 0,
						altitude: 1000,
						zoom: 16.5,
					}}
					mapType="satellite"
				/>
			</View>
		</Container>
	);
};

export default Map;

const styles = StyleSheet.create({
	container: {
		paddingHorizontal: 0,
		justifyContent: 'space-between',
	},
	title: {
		alignSelf: 'flex-start',
		marginLeft: 20,
		marginBottom: 20,
	},
	mapContainer: {
		flex: 1,
		width: '100%',
		height: '100%',
		justifyContent: 'flex-end',
		alignItems: 'center',
		overflow: 'hidden',
		borderTopLeftRadius: 50,
		borderTopRightRadius: 50,
	},
	searchFieldContainer: {
		position: 'absolute',
		top: 10,
		left: 0,
		right: 0,
		zIndex: 1,
		width: '100%',
		justifyContent: 'center',
		alignItems: 'center',
	},
	searchField: {
		backgroundColor: 'white',
		width: '90%',
		height: 70,
		borderRadius: 100,
		justifyContent: 'center',
		alignItems: 'center',
		flexDirection: 'row',
	},
	searchInput: {
		width: '80%',
		height: '100%',
		padding: 15,
	},
	searchButton: {
		width: 50,
		height: 50,
		borderRadius: 100,
		backgroundColor: Colors.primary.regular,
		justifyContent: 'center',
		alignItems: 'center',
		marginLeft: 10,
	},
	gradient: {
		width: '100%',
		height: '100%',
		position: 'absolute',
		top: 0,
		left: 0,
		right: 0,
		bottom: 0,
	},
	map: {
		width: '100%',
		height: '100%',
		overflow: 'hidden',
	},
});
