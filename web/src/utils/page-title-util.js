import { PROJECT_NAME } from '@/config/setting';

/**
 * Set browser title
 * @param title Title
 */
export function setPageTitle(title) {
  const names = [];
  if (title) {
    names.push(title);
  }
  if (PROJECT_NAME) {
    names.push(PROJECT_NAME);
  }
  document.title = names.join(' - ');
}
