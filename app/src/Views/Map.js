import {
	Dimensions,
	StyleSheet,
	Text,
	TextInput,
	TouchableOpacity,
	View,
	Animated,
} from 'react-native';
import React, { useState, useRef } from 'react';
import Container from '../Components/Container';
import MapView, { Marker } from 'react-native-maps';
import Title from '../Components/Title';
import { AntDesign } from '@expo/vector-icons';
import { Colors, Fonts } from '../style';
import { Buildings } from '../helpers/Constants';
import { t } from '../locale/translate';

const SEARCH_RESULTS_HEIGHT = 120;

const Map = () => {
	const [search, setSearch] = useState('');
	const [searchFound, setSearchFound] = useState(false);
	const [searchResults, setSearchResults] = useState('');

	const expandAnimation = useRef(new Animated.Value(0)).current;

	const onSearch = (text) => {
		const searchString = text.toUpperCase();

		const regex = /^[A-Z][1-5][0-9]{2,4}$/;

		// if empty close search
		if (searchString.length === 0) {
			setSearchFound(false);
			setSearchResults('');
			Animated.timing(expandAnimation, {
				toValue: 0,
				duration: 500,
				useNativeDriver: false,
			}).start();
			return;
		}

		Animated.timing(expandAnimation, {
			toValue: SEARCH_RESULTS_HEIGHT,
			duration: 500,
			useNativeDriver: false,
		}).start();

		if (regex.test(searchString)) {
			const building = searchString[0];
			const floor = Number(searchString[1]);

			if (!(building.toLowerCase() in Buildings)) {
				setSearchResults(t('map_not_found'));
				setSearchFound(false);
				return;
			} else if (floor > Buildings[building.toLowerCase()].floors) {
				setSearchFound(false);
				setSearchResults(t('map_not_found'));
				return;
			} else {
				setSearchFound(true);
				setSearchResults(
					`${t('go_to_building')} ${building}, ${t('room_in')} ${floor}.`
				);
			}
		} else {
			// show error message
			setSearchFound(false);
			setSearchResults(t('map_not_found'));
			return;
		}
		setSearch(text);
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
			<Title style={styles.title}>{t('campus_map')}</Title>
			<View style={styles.mapContainer}>
				<View style={styles.searchFieldContainer}>
					<View style={styles.searchField}>
						<TextInput
							style={styles.searchInput}
							onChangeText={onSearch}
							placeholder={t('search_ph')}
						/>
						<TouchableOpacity style={styles.searchButton}>
							<AntDesign name="search1" size={24} color={Colors.snowWhite} />
						</TouchableOpacity>
					</View>
					<Animated.View
						style={[styles.searchResults, { height: expandAnimation }]}
					>
						{searchFound && <Text style={styles.searchTitle}>{search}</Text>}
						<Text style={styles.searchResText}>{searchResults}</Text>
					</Animated.View>
				</View>
				<MapView
					style={styles.map}
					initialRegion={{
						latitude: 56.182252,
						longitude: 15.591309,
						latitudeDelta: 0.0922,
						longitudeDelta: 0.0421,
					}}
					initialCamera={{
						center: {
							latitude: 56.182252,
							longitude: 15.591309,
						},
						pitch: 0,
						heading: 0,
						altitude: 1000,
						zoom: 16.5,
					}}
					mapType="satellite"
				>
					{Object.keys(Buildings).map((i) => (
						<Marker
							key={i}
							coordinate={{
								latitude: Buildings[i].latitude,
								longitude: Buildings[i].longitude,
							}}
						>
							<CustomMarker title={i.toUpperCase()} />
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
		borderBottomEndRadius: 25,
		borderBottomStartRadius: 25,
		marginTop: -10,
		zIndex: -1,
		justifyContent: 'center',
		alignItems: 'center',
	},
	searchTitle: {
		fontFamily: Fonts.Inter_Bold,
		fontSize: 20,
		marginTop: 10,
	},
	searchResText: {
		padding: 10,
	},
	map: {
		width: '100%',
		height: '100%',
		overflow: 'hidden',
	},
});
