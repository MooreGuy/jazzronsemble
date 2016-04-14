import {Injectable} from 'angular2/core';
import {CONCERTS, EMPTY_CONCERT} from './mock-concerts';
import {Concert} from './concert';
import {Http, Response} from 'angular2/http';

@Injectable()	
export class ConcertService {
	// Return all mock cocerts.
	getConcerts() {
		this.http.get(this._concertsUrl).map(res =>
			<Concert[]> res.json().data).catch(this.handleError)	

		return Promise.resolve(CONCERTS);
	}

	private handleError(error: Response) {
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
