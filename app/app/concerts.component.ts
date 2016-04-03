import {Component, OnInit} from 'angular2/core';
import {Router} from 'angular2/router';

import {Concert} from './Concert';
import {ConcertDetailComponent} from './concert-detail.component';
import {CONCERTS} from './mock-concerts';
import {ConcertService} from './concert.service';

export class RonName {
	id: number;
	name: string;
	upvotes: number;
	downvotes: number;
}

@Component({
    selector: 'concerts',
	templateUrl: 'app/concerts.component.html',
	styleUrls: ['app/concerts.component.css'],
	directives: [ConcertDetailComponent]
})

export class ConcertsComponent implements OnInit {
	concerts: Concert[];

	constructor(
		private _concertService: ConcertService,
		private _router: Router
	) { }

	getConcerts() {
		this._concertService.getConcerts().then(
			concerts => this.concerts = concerts);
	}

	ngOnInit() {
		this.getConcerts();
	}

	selectedConcert: Concert;
	onSelect(concert: Concert) {
		this.selectedConcert = concert;
	}

	gotoDetail() {
		this._router.navigate(['ConcertDetail', {id: this.selectedConcert.id}]);
	}
}
