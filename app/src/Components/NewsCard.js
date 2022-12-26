import { View, StyleSheet, Text, SafeAreaView, Image} from "react-native";
import React, { startTransition } from "react";

const NewsCard = () => {
    return (
        <SafeAreaView style={styles.container}>
          <View style={{padding: 20}}>  
          {/* Title */}
            <Text style={styles.title}> Text from the Database </Text>

          {/* Description */}
            <Text style={styles.description}> This is the description </Text>

            <View style={styles.data}>
              <Text style={styles.date}> 2021-05-05 </Text>
            </View> 
            </View>

        </SafeAreaView>

    );
}

export default NewsCard;

const styles = StyleSheet.create({
    container: {
      width : '98%',
      alignSelf : 'center',
      borderRadius : 20,
      shadowOpacity : 0.5,
      shadowColor : '#000',
      shadowOffset : {width : 5, height : 5},
      backgroundColor : '#fff',
      marginTop : 20
    },
    title : {
      fontSize : 18,
      fontWeight : '600',
      merginTop : 10
    },
    description : {
      fontSize : 16,
      fontWeight : '400',
      merginTop : 10
    },
    data: {
      flexDirection : 'row',
      justifyContent : 'space-between',
      merginTop : 10 
    },
    date : {
      fontWeight : 'bold'
    }

  })