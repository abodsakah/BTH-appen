import { StyleSheet, View, Text, Settings } from 'react-native'
import React from 'react'
import Container from '../Components/Container'
import { Ionicons } from '@expo/vector-icons';
import { Colors, Fonts } from '../style'
import { AntDesign } from '@expo/vector-icons'; 



const Profile = () => {
  return (
    <Container>
      <Text  style={styles.text}>Profile</Text>
      <ProfilePage />
      <Text  style={styles.text}>More</Text>
      <Setting />
      <Language/>
      <About/>
    </Container>
  )
}
const ProfilePage = () => {
    return (
        <View style={styles.lista}>
            <Ionicons name="md-person" size={30}></Ionicons>
            <Text >Student Name</Text>
            <AntDesign name="right" size={30} color="black" />
        </View>
    )
}

const Setting = () => {
    return (
        <View style={styles.lista}>
            <Ionicons name="settings" size={30} color="black" />
            <Text >Settings</Text>
            <AntDesign name="right" size={30} color="black" />
        </View>
    )
}

const Language = () => {
    return (
        <View style={styles.lista}>
            <Ionicons name="language" size={30} color="black" />
            <Text >Language</Text>
            <AntDesign name="right" size={30} color="black" />
        </View>
    )
}

const About = () => {
    return (
        <View style={styles.lista}>
            <Ionicons name="information-circle-outline" size={30} color="black" />
            <Text style={{styles:'fontSize: Fonts.size.h4'}} >About this app</Text>
            <AntDesign name="right" size={30} color="black" />
        </View>
    )
}

export default Profile

const styles = StyleSheet.create({
    lista:{
        borderRadius: 10,
        padding:15,
        elevation: 2,
        backgroundColor: Colors.snowWhite,
        width: "100%",
        margin:5,
        alignSelf: 'flex-start', 
        flexDirection: "row",
        justifyContent: 'space-between',
        
    },
    text:{
        fontSize: Fonts.size.h2,
        alignSelf: 'flex-start', 
        fontFamily: Fonts.Inter_Bold,
        padding:5,
        margin: 3
    },
    
})