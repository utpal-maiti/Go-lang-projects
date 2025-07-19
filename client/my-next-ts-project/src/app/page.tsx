'use client';
import Image from 'next/image';
import styles from './page.module.css';
import { useEffect, useState } from 'react';

export default function Home() {
	const [status, setStatus] = useState('');

	useEffect(() => {
		fetch('http://localhost:8080/api/v1/posts')
			.then((res) => res.json())
			.then((data) => setStatus(data.status));
	}, []);

	return <h1>Status from Go API: {status}</h1>;
}
