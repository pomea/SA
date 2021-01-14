import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import foodmenu from './components/Menu';
 
export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', WelcomePage);
    router.registerRoute('/foodmenu', foodmenu);

  },
});
