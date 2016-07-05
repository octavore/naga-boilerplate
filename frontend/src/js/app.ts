let Component = {
  controller: class ComponentController { },
  view: () => m('div', 'hello world!')
};

export let routes: Mithril.Routes = {
  '/': Component
};

document.addEventListener('DOMContentLoaded', () => {
  let root = document.getElementById('app');
  m.route.mode = 'pathname';
  m.route(root, '/', routes);
});
