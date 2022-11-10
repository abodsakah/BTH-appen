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
import MapView, { Marker } from 'react-native-maps';
import Title from '../Components/Title';
import { AntDesign } from '@expo/vector-icons';
import { Colors, Fonts } from '../style';
import { Buildings } from '../helpers/Constants';

const buildingLocations = [
	{
		name: 'J',
		latitude: 56.182806,
		longitude: 15.59036,
	},
	{
		name: 'H',
		latitude: 56.182262,
		longitude: 15.590875,
	},
	{
		name: 'A',
		latitude: 56.181928,
		longitude: 15.590618,
	},
	{
		name: 'G',
		latitude: 56.181886,
		longitude: 15.591348,
	},
	{
		name: 'D',
		latitude: 56.181607,
		longitude: 15.592416,
	},
	{
		name: 'C',
		latitude: 56.181235,
		longitude: 15.592357,
	},
];

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
			}
		} else {
			// show error message
		}
	};

	const CustomMarker = ({ title }) => {
		return (
			<View style={styles.markerOuter}>
				<View style={styles.markerInner}>
					<Text style={styles.markerText}>{title}</Text>
				</View>
			</View>
		);
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
					<View style={styles.searchResults}></View>
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
				>
					{buildingLocations.map((building, index) => (
						<Marker
							coordinate={{
								latitude: building.latitude,
								longitude: building.longitude,
							}}
						>
							<CustomMarker title={building.name} />
						</Marker>
					))}
				</MapView>
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
	markerOuter: {
		backgroundColor: Colors.primary.light,
		width: 40,
		height: 40,
		borderRadius: 100,
		justifyContent: 'center',
		alignItems: 'center',
	},
	markerInner: {
		backgroundColor: Colors.primary.regular,
		width: 33,
		height: 33,
		borderRadius: 100,
		justifyContent: 'center',
		alignItems: 'center',
	},
	markerText: {
		fontFamily: Fonts.Inter_Bold,
		fontSize: 16,
		color: Colors.snowWhite,
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
	searchResults: {
		backgroundColor: Colors.snowWhite,
		width: '80%',
		height: 150,
		borderBottomEndRadius: 25,
		borderBottomStartRadius: 25,
		marginTop: -10,
	},
	map: {
		width: '100%',
		height: '100%',
		overflow: 'hidden',
	},
});
