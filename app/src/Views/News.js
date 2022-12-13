/*
TODO:

- Create layers for each news title 

- Separate the layers with some spaces

- Sync with the database

- Link to each news layer (Click on the link)
*/


import React, { Component } from 'react';
import { StyleSheet, Text, View, Image, ScrollView, TouchableOpacity } from 'react-native';
import { Container, Header, Content, Card, CardItem, Thumbnail, Button, Icon, Left, Body, Right } from 'native-base';
import { Actions } from 'react-native-router-flux';

const News = () => {
        return (
        <ScrollView>
            <Container>
            <Content>
                <Card>
                <CardItem>
                    <Left>
                    <Thumbnail source={{uri: 'Image URL'}} />
                    <Body>
                        <Text>News Title</Text>
                        <Text note>News Date</Text>
                    </Body>
                    </Left>
                </CardItem>
                <CardItem cardBody>
                    <Image source={{uri: 'Image URL'}} style={{height: 200, width: null, flex: 1}}/>
                </CardItem>
                <CardItem>
                    <Left>
                    <Button transparent>
                        <Icon active name="thumbs-up" />
                        <Text>12 Likes</Text>
                    </Button>
                    </Left>
                    <Body>
                    <Button transparent>
                        <Icon active name="chatbubbles" />
                        <Text>4 Comments</Text>
                    </Button>
                    </Body>
                    <Right>
                    <Text>11h ago</Text>
                    </Right>
                </CardItem>
                </Card>
            </Content>
            </Container>
        </ScrollView>
        );
    }

export default News;