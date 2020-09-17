import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Tesla } from '../model/tesla';

@Injectable({
  providedIn: 'root'
})
export class TeslaApiService {

  constructor(private http: HttpClient) { }

  getAllTesla(): Promise<Tesla[]> {
    return this.http.get<Tesla[]>('tesla/all').toPromise();
  }
}
