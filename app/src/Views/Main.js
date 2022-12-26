import {
	Linking,
	RefreshControl,
	ScrollView,
	StyleSheet,
	Text,
	View,
} from 'react-native';
import React, { useState, useEffect } from 'react';
import Container from '../Components/Container';
import NewsCard from '../Components/NewsCard';
import { fetchNews } from '../helpers/APIManager';
import { t } from '../locale/translate';
import { Fonts } from '../style';

const Main = () => {
	const [news, setNews] = useState([]);
	const [refreshing, setRefreshing] = useState(false);

	const getNews = async () => {
		const news = await fetchNews();
		const newsData = news?.data;

		newsData?.map((item) => {
			const date = item.date.split('T')[0];
			item.date = date;
		});

		setNews(newsData);
	};

	const refresh = async () => {
		setRefreshing(true);
		await getNews();
		setRefreshing(false);
	};

	const goToLink = (link) => {
		Linking.openURL(link);
	};

	useEffect(() => {
		getNews();
	}, []);

	return (
		<Container>
			<Text style={styles.title}>{t('news')}</Text>
			<ScrollView
				style={styles.scrollView}
				contentContainerStyle={styles.scrollViewContent}
				refreshControl={
					<RefreshControl refreshing={refreshing} onRefresh={refresh} />
				}
			>
				{news?.length > 0 ? (
					<>
						{news.map((item, index) => (
							<NewsCard
								key={index}
								title={item.title}
								description={item.description}
								date={item.date}
								onPress={() => goToLink(item.link)}
							/>
						))}
					</>
				) : (
					<Text>{t('noNews')}</Text>
				)}
			</ScrollView>
		</Container>
	);
};

export default Main;

const styles = StyleSheet.create({
	title: {
		fontFamily: Fonts.Inter_Bold,
		fontSize: Fonts.size.h1,
		marginTop: 10,
	},
	scrollView: {
		flex: 1,
		width: '100%',
	},
	scrollViewContent: {
		alignItems: 'center',
		paddingVertical: 20,
		paddingHorizontal: 10,
	},
});
