'use client';
import Image from 'next/image';
import styles from './page.module.css';
import { useEffect, useState } from 'react';
import DownloadButton from '../components/DownloadButton';

export default function Home() {
	const [status, setStatus] = useState('');

	useEffect(() => {
		fetch('http://localhost:8080/api/v1/posts')
			.then((res) => {

				for (const [key, value] of res.headers.entries()) {
        			console.log(`${key}: ${value}`);	
				}
				return res.json()
				return res;

			})
			.then((data) => {
				console.log(data)
				setStatus(data.status)
			});
	}, []);

	return (
		<>
			<h1> Status from Go API: {status}</h1>
			<DownloadButton></DownloadButton>
		</>
	);
}
