import { Component, OnInit } from '@angular/core';
import {TeslaApiService} from '../../service/tesla-api.service';
import {Tesla} from '../../model/tesla';

@Component({
  selector: 'app-tesla-cards',
  templateUrl: './tesla-cards.component.html',
  styleUrls: ['./tesla-cards.component.scss']
})
export class TeslaCardsComponent implements OnInit {

  teslas: Tesla[];

  constructor(private teslaApiService: TeslaApiService) { }

  ngOnInit(): void {
    this.teslaApiService.getAllTesla().then(value => {
      this.teslas = value;
      console.log(value);
    });
  }

}
