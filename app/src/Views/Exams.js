import {
	RefreshControl,
	ScrollView,
	StyleSheet,
	Text,
	View,
} from 'react-native';
import React, { useEffect, useState } from 'react';
import Container from '../Components/Container';
import Button from '../Components/Button';
import { Colors, Fonts } from '../style';
import Exam from '../Components/Exam';
import { t } from '../locale/translate';
import {
	listAllExams,
	listUserExams,
	registerExam,
	unregisterExam,
} from '../helpers/APIManager';

const Exams = () => {
	const [registeredExams, setRegisteredExams] = useState([]);
	const [commingExams, setCommingExams] = useState([]);
	const [refreshing, setRefreshing] = useState(false);

	const getExams = async () => {
		// get the registered exams
		let re = await listUserExams();
		if (re?.status === 200) {
			setRegisteredExams(re?.data);
		}

		let ce = await listAllExams();
		if (ce?.status === 200) {
			// remove the re from ce
			let exams = ce?.data.filter((e) => {
				return !re?.data.find((r) => r.ID === e.ID);
			});
			setCommingExams(exams);
		}
	};

	const refresh = async () => {
		setRefreshing(true);
		await getExams();
		setRefreshing(false);
	};

	const HandleRegisterExam = async (id) => {
		let res = await registerExam(id);
		if (res.status === 200) {
			await getExams();
		}
	};

	const HandleUnregisterExam = async (id) => {
		let res = await unregisterExam(id);
		if (res.status === 200) {
			await getExams();
		}
	};

	useEffect(() => {
		getExams();
	}, []);

	return (
		<Container style={styles.container}>
			<Text style={styles.title}>{t('exams')}</Text>
			<ScrollView
				style={styles.scrollContainer}
				contentContainerStyle={styles.scrollContentStyle}
				refreshControl={
					<RefreshControl
						refreshing={refreshing}
						onRefresh={refresh}
						colors={[Colors.primary.regular]}
					/>
				}
			>
				{registeredExams.length > 0 && (
					<>
						<Text style={styles.heading}>{t('registered_exams')}</Text>
						{registeredExams.map((exam, i) => (
							<Exam
								key={i}
								HandleRegisterExam={() => HandleRegisterExam(exam.ID)}
								HandleUnregisterExam={() => HandleUnregisterExam(exam.ID)}
								code={exam.course_code}
								name={exam.name}
								date={new Date(exam.start_date).toLocaleDateString()}
								time={new Date(exam.start_date)
									.toLocaleTimeString()
									.slice(0, 5)}
								room={exam?.room || 'Not set'}
								registered
							/>
						))}
					</>
				)}
				{commingExams.length > 0 && (
					<>
						<Text style={styles.heading}>{t('coming_exams')}</Text>
						{commingExams.map((exam, i) => (
							<Exam
								key={i}
								code={exam.course_code}
								name={exam.name}
								date={new Date(exam.start_date).toLocaleDateString()}
								time={new Date(exam.start_date)
									.toLocaleTimeString()
									.slice(0, 5)}
								room={exam?.room || 'Not set'}
								HandleRegisterExam={() => HandleRegisterExam(exam.ID)}
								HandleUnregisterExam={() => HandleUnregisterExam(exam.ID)}
							/>
						))}
					</>
				)}
			</ScrollView>
		</Container>
	);
};

export default Exams;

const styles = StyleSheet.create({
	container: {
		paddingHorizontal: 0,
	},
	title: {
		fontSize: Fonts.size.h1,
		alignSelf: 'flex-start',
		fontFamily: Fonts.Inter_Bold,
		marginLeft: 10,
	},
	scrollContainer: {
		width: '100%',
	},
	scrollContentStyle: {
		paddingHorizontal: 15,
		alignItems: 'center',
		paddingBottom: 20,
	},
	heading: {
		fontSize: Fonts.size.h3,
		margin: 5,
		alignSelf: 'flex-start',
	},
});
