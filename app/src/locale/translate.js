import translation from './translation.json';

export let selectedLanguage = 'en';

export function setLanguage(languageCode) {
	switch (languageCode) {
		case 'en':
		case 'en-US':
		case 'en-GB':
		case 'en-AU':
		case 'en-CA':
			selectedLanguage = 'en';
			break;
		case 'se':
		case 'se-SE':
		case 'se-FI':
		case 'sv':
		case 'sv-sv':
		case 'sv-SE':
		case 'sv-FI':
		case 'sv-fi':
			selectedLanguage = 'sv-sv';
			break;
		case 'ar':
		case 'ar-SA':
		case 'ar-EG':
		case 'ar-IQ':
		case 'ar-JO':
		case 'ar-KW':
		case 'ar-LB':
		case 'ar-LY':
		case 'ar-MA':
		case 'ar-OM':
		case 'ar-QA':
		case 'ar-SA':
		case 'ar-SY':
		case 'ar-TN':
		case 'ar-AE':
		case 'ar-YE':
			selectedLanguage = 'ar';
			break;
		default:
			selectedLanguage = 'en';
			break;
	}
}

export function t(key) {
	return translation[key] ? translation[key][selectedLanguage] : key;
}
