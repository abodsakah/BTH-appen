import { StyleSheet, Text, View } from 'react-native'
import React from 'react'
import Container from '../Components/Container'
import Button from '../Components/Button'
import { Colors, Fonts } from '../style'

const Exams = () => {
  return (
    <Container>
      <Text style={styles.code}>Exams</Text>
      <Text style={styles.textStyle}>Registered exams</Text>
      <Register />
      <Register2 />
      <Text style={styles.textStyle}>Comming exams</Text>
      <Comming />
    </Container>
  )
}

const Register = () => {
    return (
    <View style={styles.examContainer}>
        <Text style={styles.code}>PA1469</Text>
        <Text>Applikation Utveckling</Text>
        <Text>Date: 2023-05-31</Text>
        <Text>Time: 13:00</Text>
        <Text>Room: J1240</Text>
        <Button type="danger" style={styles.buttonStyle}>Unregister</Button>
    </View>
    )
}

const Register2 = () => {
    return (
    <View style={styles.examContainer}>
        <Text style={styles.code}>MA1476</Text>
        <Text>Matimatik introduktion</Text>
        <Text>Date: 2023-05-31</Text>
        <Text>Time: 13:00</Text>
        <Text>Room: J1240</Text>
        <Button type="danger" style={styles.buttonStyle}>Unregister</Button>
    </View>
    )
}

const Comming = () => {
    return (
    <View style={styles.examContainer}>
        <Text style={styles.code}>
            DV1628</Text>
        <Text style={{flexDirection: 'row'}}>Datorteknik</Text>
        <Text style={{flexDirection: 'row'}}>Date: 2023-05-31</Text>
        <Text>Time: 13:00</Text>
        <Text>Room: J1240</Text>
        <Button style={styles.buttonStyle}>Register</Button>
    </View>
    )
}

export default Exams

const styles = StyleSheet.create({
    examContainer:{
        padding: 10,
        elevation: 10,
        width: "100%",
        backgroundColor: Colors.snowWhite,
        justifyContent: 'flex-start',
        borderRadius: 15,
        margin:5
    },
    code:{
        fontSize: Fonts.size.h3,
        alignSelf: 'flex-start', 
        fontFamily: Fonts.Inter_Bold,
    },
    buttonStyle:{
        width: "50%",
        alignSelf: "flex-end"
    }, 
    textStyle:{
        fontSize: Fonts.size.h3,
        borderRadius: 15,
        margin:5,
        alignSelf: 'flex-start', 
    }
})