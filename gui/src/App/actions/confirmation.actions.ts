/**
 * Created by justin on 8/12/16.
 */

export const OPEN_CONFIRMATION = 'OPEN_CONFIRMATION';
export const CLOSE_CONFIRMATION = 'CLOSE_CONFIRMATION';

export function openConfirmation(title, text, onYes, onNo) {
  return {
    type: OPEN_CONFIRMATION,
    text,
    title,
    onYes,
    onNo,
    isOpen: true
  };
}

export function closeConfirmation() {
  return {
    type: CLOSE_CONFIRMATION
  };
}
