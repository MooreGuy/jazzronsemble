import {Concert} from './concert';
import {ConcertService} from './concert.service';

import {RouteParams} from 'angular2/router';
import { Component, Input, OnInit } from 'angular2/core';

@Component({
	selector: 'concert-detail',
	providers: [ConcertService],
	templateUrl: 'app/concert-detail.component.html',
	styleUrls: ['app/concert-detail.component.css']
})

export class ConcertDetailComponent implements OnInit {
	@Input()
	concert: Concert;

	constructor(
		private _concertService: ConcertService,
		private _routeParams: RouteParams) {}

	ngOnInit() {
		let id = +this._routeParams.get('id');
		this._concertService.getConcert(id).then(
			concert => this.concert = concert);
	}
}
