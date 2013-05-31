def wait_and_kill(pids)
  trap("INT") {
    pids.each { |pid| Process.kill(3, pid) rescue Errno::ESRCH }
    exit 0
  }

  pids.each { |pid| Process.wait(pid) }
end

desc "Compile site"
task :compile do
  jekyllPid = Process.spawn("jekyll build")
  compassPid = Process.spawn("compass compile ./_sass -e #{ENV['ENV'] || 'development'}")
  wait_and_kill [jekyllPid, compassPid]
end

desc "Watch"
task :watch do
  jekyllPid = Process.spawn("jekyll build --watch")
  compassPid = Process.spawn("compass watch ./_sass -e #{ENV['ENV'] || 'development'}")
  wait_and_kill [jekyllPid, compassPid]
end

desc "Preview"
task :preview do
  jekyllPid = Process.spawn("jekyll serve --watch")
  compassPid = Process.spawn("compass watch ./_sass -e #{ENV['ENV'] || 'development'}")
  wait_and_kill [jekyllPid, compassPid]
end

desc "Deploy to heroku"
task :deploy do
  if `git remote -v|grep "heroku"`.empty?
    `git remote add heroku git@heroku.com:unix-foo-6w.git`
  end
  `git push heroku master`
end
