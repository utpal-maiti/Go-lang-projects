import { downloadFile } from '@/utils/downloadFile';

const DownloadButton = () => (
	<>
		<button
			onClick={() =>
				downloadFile('http://localhost:8080/download', 'sample.pdf')
			}
		>
			Download Sample PDF
		</button>
	</>
);

export default DownloadButton;
