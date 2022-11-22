import { StyleSheet, View, Text, Settings } from 'react-native'
import React from 'react'
import Container from '../Components/Container';
import { Colors, Fonts } from '../style';
import { AntDesign, Ionicons, Entypo } from '@expo/vector-icons';
import OptionContainer from '../Components/OptionContainer';
import { t } from '../locale/translate';

const Profile = ({ navigation }) => {
	const navigateToLanguages = () => {
		navigation.navigate('Languages');
	};

	return (
		<Container style={styles.container}>
			<Text style={styles.heading}>Profile</Text>
			<OptionContainer
				text="Student Name"
				Icon={() => <Ionicons name="md-person" size={30}></Ionicons>}
			/>
			<Text style={styles.heading}>More</Text>
			<OptionContainer
				text={t('settings')}
				Icon={() => (
					<AntDesign name="setting" size={30} color={Colors.primary.regular} />
				)}
			/>
			<OptionContainer
				onPress={navigateToLanguages}
				text={t('language')}
				Icon={() => (
					<Entypo name="language" size={30} color={Colors.primary.regular} />
				)}
			/>
			<OptionContainer
				text={t('about')}
				Icon={() => (
					<Ionicons
						name="md-information-circle"
						size={30}
						color={Colors.primary.regular}
					/>
				)}
			/>
		</Container>
	);
};
const ProfilePage = () => {
	return (
		<View style={styles.lista}>
			<Ionicons name="md-person" size={30}></Ionicons>
			<Text>Student Name</Text>
			<AntDesign name="right" size={30} color="black" />
		</View>
	);
};

const Setting = () => {
	return (
		<View style={styles.lista}>
			<Ionicons name="settings" size={30} color="black" />
			<Text>Settings</Text>
			<AntDesign name="right" size={30} color="black" />
		</View>
	);
};

const Language = () => {
	return (
		<View style={styles.lista}>
			<Ionicons name="language" size={30} color="black" />
			<Text>Language</Text>
			<AntDesign name="right" size={30} color="black" />
		</View>
	);
};

const About = () => {
	return (
		<View style={styles.lista}>
			<Ionicons name="information-circle-outline" size={30} color="black" />
			<Text style={{ styles: 'fontSize: Fonts.size.h4' }}>About this app</Text>
			<AntDesign name="right" size={30} color="black" />
		</View>
	);
};

export default Profile;

const styles = StyleSheet.create({
	container: {
		justifyContent: 'flex-start',
	},
	heading: {
		fontSize: Fonts.size.h2,
		alignSelf: 'flex-start',
		fontFamily: Fonts.Inter_Bold,
		padding: 5,
		margin: 3,
	},
});