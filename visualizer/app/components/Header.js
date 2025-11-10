import { VscJson } from 'react-icons/vsc';

export default function Header() {
  return (
    <header className="shadow-md bg-bg-dark relative after:absolute after:bottom-0 after:left-0 after:right-0 after:h-[2px] after:bg-gradient-to-r after:from-ct-yellow after:via-ct-green after:to-ct-purple">
      <div className="container mx-auto py-3">
        <h1 className="text-2xl font-bold m-0 text-ct-yellow flex items-center gap-2">
          <VscJson className="text-ct-green" />
          deepspec
        </h1>
      </div>
    </header>
  );
}
