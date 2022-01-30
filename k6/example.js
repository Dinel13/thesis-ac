import http from 'k6/http';
import { sleep } from 'k6';

// default option tidak pelru ketik di terminal
export const options = {
   vus: 10,
   duration: '30s',
 };

export default function () {
  http.get('https://test.k6.io');
  sleep(1);
}
