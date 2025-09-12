import { setupWorker } from 'msw/browser';

const worker = setupWorker();
// http.get('/admin/admin/member/index', ({ request, params, cookies }) => {
//   return HttpResponse.json(memberList);
// }),
// http.get('/admin/admin/memberRechargeOrder/rechargeReport', ({ request, params, cookies }) => {
//   return HttpResponse.json(rechargeReport);
// }),

await worker.start({
  onUnhandledRequest: 'bypass' // Let unhandled requests pass through
});
