import styles from './schedulii-components.module.css';

/* eslint-disable-next-line */
export interface ScheduliiComponentsProps {}

export function ScheduliiComponents(props: ScheduliiComponentsProps) {
  return (
    <div className={styles['container']}>
      <h1>Welcome to ScheduliiComponents!</h1>
    </div>
  );
}

export default ScheduliiComponents;
