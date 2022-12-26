import React, { Component } from 'react';
import { StyleSheet, Text, View, SafeAreaView } from 'react-native';
import NewsCard from '../Components/NewsCard';


const News = () => {
        return (
        <SafeAreaView style={style.container}>
            <NewsCard/>
            <NewsCard/>
            <NewsCard/>
        </SafeAreaView>
        );
    }

export default News;