import {Component, OnInit} from 'angular2/core';
import {Concert} from './concert';
import {ConcertService} from './concert.service';

import {Router} from 'angular2/router';

@Component({
	selector: 'dashboard',
	templateUrl: 'app/dashboard.component.html',
	styleUrls: ['app/dashboard.component.css'],
})

export class DashboardComponent implements OnInit {
	concerts: Concert[] = [];
	upcomingConcert : Concert;
	emptyConcert: Concert;

	constructor(
		private _concertService: ConcertService,
		private _router: Router
	) { }

	ngOnInit() {
		this._concertService.getEmptyConcert().then(
			emptyConcert => this.emptyConcert = emptyConcert);
		this._concertService.getConcerts().
			then(concerts => this.concerts = concerts);
		this.upcomingConcert = (this.concerts.length > 0) ? this.concerts[0] :
			this.emptyConcert;
	}
}
