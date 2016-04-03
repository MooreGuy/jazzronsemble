import {Injectable} from 'angular2/core';
import {CONCERTS, EMPTY_CONCERT} from './mock-concerts';

@Injectable()	
export class ConcertService {
	// Return all mock cocerts.
	getConcerts() {
		return Promise.resolve(CONCERTS);
	}

	// Empty concerts for when no concert exists, something has to be displayed
	getEmptyConcert() {
		return Promise.resolve(EMPTY_CONCERT);
	}

	// Get a single concert by ID.
	getConcert(id: number) {
		return Promise.resolve(CONCERTS).then(
			concerts => concerts.filter(concert => concert.id === id)[0]);
	}
}
