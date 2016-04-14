import {Component} from 'angular2/core';
import {ConcertService} from './concert.service';
import {ConcertsComponent} from './concerts.component';
import {DashboardComponent} from './dashboard.component';
import {ConcertDetailComponent} from './concert-detail.component';
import {AboutComponent} from './about.component';
import {ContactComponent} from './contact.component';
import {RonName} from './RonName';

import {
	RouteConfig,
	ROUTER_DIRECTIVES,
	ROUTER_PROVIDERS
} from 'angular2/router';

@Component({
	selector: 'my-app',
	templateUrl: 'static/app/app.component.html',
	styleUrls: ['static/app/app.component.css'],
	directives: [ROUTER_DIRECTIVES],
	providers: [
		ConcertService,
		ROUTER_PROVIDERS
	]
})

@RouteConfig([
	{
		path: '/concerts',
		name: 'Concerts',
		component: ConcertsComponent
	},
	{
		path: '/concerts/detail/:id',
		name: 'ConcertDetail',
		component: ConcertDetailComponent
	},
	{
		path: '/about',
		name: 'About',
		component: AboutComponent
	},
	{
		path: '/contact',
		name: 'Contact',
		component: ContactComponent

	}
])

export class AppComponent {
	// TODO: Static variables should be removed at some point.
	title = 'Cuesta College Jazz Ensemble';

	// TODO: Move this to it's own component.
	directorName: RonName = {
		id: 1,
		name: 'CD-Ron',
		upvotes: 1,
		downvotes: 0
	};

	goBack() {
		window.history.back();
	}
}
